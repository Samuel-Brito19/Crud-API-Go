package testing_test

import (
	"go-api/model"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

func TestTesting(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Testing Suite")
}

var _ = Describe("Let's test that shit",Ordered, func() {

	var (
		server *ghttp.Server
		body []*model.Product
	)

	BeforeEach(func() {
		server = ghttp.NewServer()
		Expect(server).ToNot(BeNil())
		body = []*model.Product{
			{
				ID: 3,
				Name: "macaco",
				Price: 4342,
			},
		}
		server.AppendHandlers(
			ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", "/products"),
				ghttp.RespondWithJSONEncoded(200, &body),
			),
		)		
	})


	Context("Should get the products", func() {
		It("Should return status 200", func() {
			resp, err := server.HTTPTestServer.Client().Get(server.URL() + "/products")
			Expect(err).ToNot(HaveOccurred())
			Expect(resp.Status).To(Equal("200 OK"))
		})
	})

	AfterAll(func() {
		server.Close()
	})
})