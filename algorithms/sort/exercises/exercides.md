### Which sorting algorithm should be used?

## 1 - Sort 10 schools around your house by distance:

Insertion Sort, since the number of elements to be sorted are SMALL. Space complexity of O(1)

## 2 - eBay sorts listings by the current Bid amount:

Radix or Counting Sort. These algorithms are very fast to compare a SMALL amount of INTEGER NUMBERS (and only for numbers) (number of bids is an integer)

## 3 - Sport scores on ESPN

Quick Sort. Merge Sort might lead to big space complexity.


## 4 - Massive database (can't fit all into memory) needs to sort through past year's user data

Merge Sort, because we are not sorting in memory and since the data is so big, we should worry
about the WORST CASE SCENARIO, where Merge Sort if faster than Quick Sort

## 5 - Almost sorted Udemy review data needs to update and add 2 new reviews

Insertion Sort, since the data is almost sorteda and this makes this algorithm very fast (O(n)).

## 6 - Temperature Records for the past 50 years in Canada

1)Quick Sort, since the amout of data is huge and probably would be very slow to store all those data
in memory by using Merge Sort

2)Radix Sort or Counting Sort if the Temperature Records have not decimal place (thus, these sorting algorithms are better for sorting interger numbers)

## 7 - Large user name database needs to be sorted. Data is very random.

1) If we have enought memory space => Merge Sort, since the randomness could lead to the selection of a bad pivot in Quick Sort (leads to the worst case (O(n^2)) )

## 8 - You want to teach sorting for the first time

Bubble Sort and Selection Sort algorithms.