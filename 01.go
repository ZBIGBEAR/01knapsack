package main

import "fmt"

func main(){
	var n int
	fmt.Println("请输入物品数量：")
	fmt.Scanln(&n)

	//var weights [n]int
	weights := make([]int,n)
	fmt.Println("请输入各个物品的重量：")
	for i:=0;i<n;i++{
		fmt.Scanln(&weights[i])
	}
	fmt.Println("请输入各个物品的价值：")
	//var values [n]int
	values := make([]int,n)
	for i:=0;i<n;i++{
		fmt.Scanln(&values[i])
	}
	fmt.Println("请输入背包能承载的重量：")
	var c int
	fmt.Scanln(&c)
	
	fmt.Println(n,weights,values,c)
	m := knapsack(n,weights,values,c)
	print2Int(m)
}

/*
param1:物品的数量
param2:各个物品的重量
param3:各个物品的价值
param4:背包的容量
算法说明
m[i][j]:表示面对第i个物品，背包容量为j的时候，能获得的最大价值
if j<weights[i]   //背包掏空了都放不下第i个物品
	m[i][j]=m[i-1][j]
else   //背包掏空了可以放下第i个物品
	m[i][j]=max(m[i-1][j],m[i-1][j-w[i]]+v[i])
*/
func knapsack(n int,weights []int,values []int,c int) [][]int{
	//定义二维数组
	m := make([][]int,(n+1))
	for i:=0;i<n+1;i++{
		m[i]=make([]int,(c+1))
	}
	//初始化m[1][1...c]
	for i:=1;i<=c;i++{
		if weights[0]>i{
			m[1][i]=0
		}else{
			m[1][i]=values[0]
		}
	}
	//求数组中剩余变量的值
	for i:=2;i<=n;i++{
		for j:=1;j<=c;j++{
			if j<weights[i-1]{
				m[i][j]=m[i-1][j]
			}else{
				//fmt.Println(i,j,j-weights[i-1])
				m[i][j]=max(m[i-1][j],m[i-1][j-weights[i-1]]+values[i-1])
			}
		}
	}
	return m
}

func print2Int(m [][]int){
	for i:=1;i<len(m);i++{
		for j:=1;j<len(m[i]);j++{
			fmt.Print(m[i][j],"\t")
		}
		fmt.Println()
	}
}

func max(a,b int) int {
	if a>b {
		return a
	}else{
		return b
	}
}
