package recur

/*
	큰 문제를 작은 문제로 나눠 푸는 방법

	- Props
	Stack Memory를 사용 -> 굉장히 효율적

	- Cons
	하지만 Stack Memory를 많이 사용하면 -> StackOverFlow (재귀호출이 너무 깊은 경우)
*/

func Fibo(num int) int {

	// exit
	if num <=1 {
		return num
	}

	return Fibo(num-2) + Fibo(num-1)
}

/*
	꼬리호출

	- 함수코드 제일 마지막에서 다른 함수를 호출하는 경우
	- 재귀 같은 경우에는 결국 -> 스택프레임을 사용한다
	- 꼬리재귀 형태를 사용 (마지막에 return Fn) 하는경우에는 기존에 스택프레임을 유지 하지 않아도 된다.
	- TCO가 최적화 되어있다면 -> 재귀형태이지만 스택프레임을 유지하지 않는다.
	* 스택프레임을 사용하는 재귀형태이지만, StackOverFlow가 나지 않는다 (TCO 최적화가 된다는 점에서...)
*/

// func recurTotal(num int) int {
	
	
// }