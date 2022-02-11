# 盛最多水的容器

## 题目描述

给出一个非负整数数组 a1，a2，a3，…… an，每个整数标识一个竖立在坐标轴 x 位置的一堵高度为 ai 的墙，选择两堵墙，和 x 轴构成的容器可以容纳最多的水。

**示例1：**

![](https://aliyun-lc-upload.oss-cn-hangzhou.aliyuncs.com/aliyun-lc-upload/uploads/2018/07/25/question_11.jpg)
```
输入：[1,8,6,2,5,4,8,3,7]
输出：49 
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
```

**示例2：**

```
输入：height = [1,1]
输出：1
```

**示例3：**

```
输入：height = [4,3,2,1,4]
输出：16
```

**提示：**

- n == height.length
- 2 <= n <= 105
- 0 <= height[i] <= 104

## 解题思路

对撞指针的思路。首尾分别 2 个指针，每次移动以后都分别判断长宽的乘积是否最大。