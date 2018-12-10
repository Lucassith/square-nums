package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSquareSums(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SquareSums Suite")
}
