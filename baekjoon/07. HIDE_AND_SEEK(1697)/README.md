## BFS 로 풀 경우 
## CUR -1, CUR +1, CUR *2  를 사용 
## NCUR이 0 <= NCUR <= MAXCOUNT 만큼이어야 한다.
## 그리고 이미 방문하지 않은 노드여야 한다 . 

## BFS로 푼 이유 
## 이미 방문한 x축에서는 재 방문을 안하는 이유 ???? 
## 이미 계산을 해본 Node이기 때문에

5, 17 예제 손풀이 
0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21  
          N                              k 

                 5
    4            6          10
3  [5]  8   [5]   7  12   9  11  20   
