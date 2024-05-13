// 2023/1/2,19:35
package main

func fbnArr(n int) []uint64 {
	fbnSli := make([]uint64, n)
	if n >= 1 {
		fbnSli[0] = 1
		if n >= 2 {
			fbnSli[1] = 1
			if n >= 3 {
				for i := 2; i < n; i++ {
					fbnSli[i] = fbnSli[i-1] + fbnSli[i-2]
				}
			}
		}
	}
	return fbnSli
}
func main() {

	//fmt.Println(fbnArr(2))

}
