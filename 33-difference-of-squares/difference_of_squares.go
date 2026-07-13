package differenceofsquares

func SquareOfSum(n int) int {
	var sayilar []int
	toplam1 := 0
	for i := 1; i < n+1; i++ {
		sayilar = append(sayilar, i)
	}
	for _, deger := range sayilar {
		toplam1 = toplam1 + deger
	}
	return toplam1 * toplam1

}

func SumOfSquares(n int) int {
	var sayilar []int
	toplam2 := 0
	for i := 1; i < n+1; i++ {
		sayilar = append(sayilar, i*i)
	}
	for _, deger := range sayilar {
		toplam2 = toplam2 + deger
	}
	return toplam2
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
