package iio

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Context("IIO", func() {
	Context("New", func() {
		It("should return IIO", func() {
			/* arrange/act/assert */
			Expect(New()).
				Should(Not(BeNil()))
		})
	})
	Context("Pipe", func() {
		It("should create non nil reader & writer", func() {
			/* arrange */
			objectUnderTest := New()

			/* act */
			actualReader, actualWriter := objectUnderTest.Pipe()
			actualWriter.Close()
			actualReader.Close()

			/* assert */
			Expect(actualReader).NotTo(BeNil())
			Expect(actualWriter).NotTo(BeNil())
		})
	})
})
