package integration_test

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"

	"code.cloudfoundry.org/groot/integration/cmd/foot/foot"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("pull", func() {
	var rootfsURI string

	BeforeEach(func() {
		var err error
		tempDir, err = ioutil.TempDir("", "groot-integration-tests")
		Expect(err).NotTo(HaveOccurred())
		configFilePath = filepath.Join(tempDir, "groot-config.yml")
		rootfsURI = filepath.Join(tempDir, "rootfs.tar")

		logLevel = ""
		env = []string{"FOOT_BASE_DIR=" + tempDir}
		stdout = new(bytes.Buffer)
		stderr = new(bytes.Buffer)
	})

	AfterEach(func() {
		Expect(os.RemoveAll(tempDir)).To(Succeed())
	})

	runPullCmd := func() error {
		footArgv := []string{"--config", configFilePath, "pull", rootfsURI}
		footCmd := exec.Command(footBinPath, footArgv...)
		footCmd.Stdout = io.MultiWriter(stdout, GinkgoWriter)
		footCmd.Stderr = io.MultiWriter(stderr, GinkgoWriter)
		footCmd.Env = append(os.Environ(), env...)
		return footCmd.Run()
	}

	Describe("success", func() {
		JustBeforeEach(func() {
			if configFilePath != "" {
				configYml := fmt.Sprintf(`log_level: %s`, logLevel)
				Expect(ioutil.WriteFile(configFilePath, []byte(configYml), 0600)).To(Succeed())
			}
		})

		Describe("Local images", func() {
			JustBeforeEach(func() {
				Expect(ioutil.WriteFile(rootfsURI, []byte("a-rootfs"), 0600)).To(Succeed())

				Expect(runPullCmd()).To(Succeed())
			})

			Context("when pull succeeds", func() {
				unpackIsSuccessful(runPullCmd)
			})

			It("calls driver.Unpack() with the correct stream", func() {
				var args foot.UnpackCalls
				readTestArgsFile(foot.UnpackArgsFileName, &args)
				Expect(string(args[0].LayerTarContents)).To(Equal("a-rootfs"))
			})

			Describe("subsequent invocations", func() {
				Context("when the rootfs file timestamp has changed", func() {
					It("generates a different layer ID", func() {
						var unpackArgs foot.UnpackCalls
						readTestArgsFile(foot.UnpackArgsFileName, &unpackArgs)
						firstInvocationLayerID := unpackArgs[0].ID

						now := time.Now()
						Expect(os.Chtimes(rootfsURI, now.Add(time.Hour), now.Add(time.Hour))).To(Succeed())

						Expect(runPullCmd()).To(Succeed())

						readTestArgsFile(foot.UnpackArgsFileName, &unpackArgs)
						secondInvocationLayerID := unpackArgs[1].ID

						Expect(secondInvocationLayerID).NotTo(Equal(firstInvocationLayerID))
					})
				})
			})
		})

		Describe("Remote images", func() {
			JustBeforeEach(func() {
				rootfsURI = "docker:///cfgarden/three-layers"

				Expect(runPullCmd()).To(Succeed())
			})

			Context("when pull succeeds", func() {
				unpackIsSuccessful(runPullCmd)
			})

			Context("when the image has multiple layers", func() {
				It("correctly passes parent IDs to each driver.Unpack() call", func() {
					var args foot.UnpackCalls
					readTestArgsFile(foot.UnpackArgsFileName, &args)

					chainIDs := []string{}
					for _, a := range args {
						Expect(a.ParentIDs).To(Equal(chainIDs))
						chainIDs = append(chainIDs, a.ID)
					}
				})
			})
		})
	})

	Describe("failure", func() {
		Describe("Local Images", func() {
			var createRootfsTar bool

			BeforeEach(func() {
				createRootfsTar = true
			})

			JustBeforeEach(func() {
				if createRootfsTar {
					Expect(ioutil.WriteFile(rootfsURI, []byte("a-rootfs"), 0600)).To(Succeed())
				}
			})

			whenUnpackIsUnsuccessful(runPullCmd)

			Context("when the rootfs URI is not a file", func() {
				BeforeEach(func() {
					createRootfsTar = false
				})

				It("prints an error", func() {
					Expect(stdout.String()).To(ContainSubstring(notFoundRuntimeError[runtime.GOOS]))
				})
			})
		})

		Describe("Remote Images", func() {
			BeforeEach(func() {
				rootfsURI = "docker:///cfgarden/three-layers"
			})

			whenUnpackIsUnsuccessful(runPullCmd)
		})

	})
})
