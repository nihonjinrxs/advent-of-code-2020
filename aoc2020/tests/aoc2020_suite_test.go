package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAoc2020(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AoC2020 Suite")
}
