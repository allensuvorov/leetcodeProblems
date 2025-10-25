/*
1 - loss
2 - win
any even number - wins, as it can be changed into odd, by deducting a 1
any odd number - loses, as it can only be changed into an even number
*/
func divisorGame(n int) bool {
    return n % 2 == 0
}
