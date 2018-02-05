package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
)

var _ = Describe("UserLogin", func() {
	var page *agouti.Page

	BeforeEach(func() {

		var err error
		page, err = agoutiDriver.NewPage(agouti.Browser("Chrome"))
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		Expect(page.Destroy()).To(Succeed())
	})

	It("should manage user authentication", func() {
		By("redirecting the user to the login form from the home page", func() {
			Expect(page.Navigate("chrome://extensions/")).To(Succeed())

		})

		By("allowing the user to fill out the login form and submit it", func() {

		})

		By("allowing the user to view their profile", func() {

		})

		By("allowing the user to log out", func() {

		})
	})
})
