package service

import (
	"testing"
)

func TestConvert(t *testing.T) {
	t.Run("Testing Convert : should pass", func(t *testing.T) {
		err := Convert("../../tests/config.yaml")
		if err != nil {
			t.Fatalf("Should not fail : found error %v", err)
		}
	})

	t.Run("Testing Convert : (bad config) should fail", func(t *testing.T) {
		err := Convert("../../tests/conf.yaml")
		if err == nil {
			t.Fatalf("Should fail : found error %v", err)
		}
	})

	t.Run("Testing Convert : (bad bc) should fail", func(t *testing.T) {
		err := Convert("../../tests/config-bad-bc.yaml")
		if err == nil {
			t.Fatalf("Should fail : found error %v", err)
		}
	})

}
