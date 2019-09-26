package main

import (
	"testing" // テストで使える関数・構造体が用意されているパッケージをimport

	. "github.com/smicle/gospel"
)

func TestExample(t *testing.T) {
	Describe(t, "fizzbuzz", func() {
		Context("3の倍数を渡したら、Fizzを返す", func() {
			It("3を渡したら、Fizzを返す", func() {
				Expect(fizzbuzz(3)).To(Equal, "Fizz")
			})
			It("6を渡したら、Fizzを返す", func() {
				Expect(fizzbuzz(6)).To(Equal, "Fizz")
			})
		})
		Context("5の倍数を渡したら、Buzzを返す", func() {
			It("5を渡したら、Buzzを返す", func() {
				Expect(fizzbuzz(5)).To(Equal, "Buzz")
			})
			It("10を渡したら、Buzzを返す", func() {
				Expect(fizzbuzz(10)).To(Equal, "Buzz")
			})
		})
		Context("3と5の倍数を渡したら、FizzBuzzを返す", func() {
			It("15を渡したら、FizzBuzzを返す", func() {
				Expect(fizzbuzz(15)).To(Equal, "FizzBuzz")
			})
			It("30を渡したら、FizzBuzzを返す", func() {
				Expect(fizzbuzz(30)).To(Equal, "FizzBuzz")
			})
		})
		Context("それ以外の数字を渡したら、数字をそのまま返す", func() {
			It("1を渡したら、1を返す", func() {
				Expect(fizzbuzz(1)).To(Equal, "1")
			})
			It("2を渡したら、2を返す", func() {
				Expect(fizzbuzz(2)).To(Equal, "2")
			})
		})
	})
}
