package asm_test

import (
	"os"
	"testing"

	"github.com/ayoul3/asm-env/asm"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"
)

func TestAWS(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecsWithDefaultAndCustomReporters(t, "AWS/SM", []Reporter{reporters.NewJUnitReporter("test_report-aws-sm.xml")})
}

var _ = Describe("SM", func() {
	Describe("GetSecret", func() {
		BeforeEach(func() {
			os.Setenv("AWS_REGION", "eu-west-1")
		})
		AfterEach(func() {
			os.Setenv("AWS_REGION", "")
		})
		Context("When the client fails", func() {
			It("should return an error", func() {
				client, _ := asm.NewClient(&asm.MockClient{GetSecretShouldFail: true})
				_, err := client.GetSecret("arn:aws:secretsmanager:eu-west-1:123456789123:secret/key1")
				Expect(err).To(HaveOccurred())
			})
		})
		Context("When the keys contains an index", func() {
			It("it should return the secret", func() {
				client, _ := asm.NewClient(&asm.MockClient{})
				secret, err := client.GetSecret("arn:aws:secretsmanager:eu-west-1:123456789123:secret/key1#index")
				Expect(err).ToNot(HaveOccurred())
				Expect(secret).To(Equal(asm.MockSecretValue))
			})
		})
		Context("When the call succeeds", func() {
			It("it should return the secret", func() {
				client, _ := asm.NewClient(&asm.MockClient{})
				secret, err := client.GetSecret("arn:aws:secretsmanager:eu-west-1:123456789123:secret/key1")
				Expect(err).ToNot(HaveOccurred())
				Expect(secret).To(Equal(asm.MockSecretValue))
			})
		})
	})
})
