package main

func part2(input string) int {
	algo, image := prepareData(input)

	return len(enhanceImage(image, algo, 50))
}
