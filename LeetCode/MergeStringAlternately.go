/*
https://leetcode.com/problems/merge-strings-alternately/
*/

func mergeAlternately(word1 string, word2 string) string {
    merged := make([]byte, 0, len(word1)+len(word2))
    for i:=0;i<len(word1) || i<len(word2);i++{
        if i<len(word1){
        merged=append(merged,word1[i])
        }
        if i<len(word2){
        merged=append(merged,word2[i])
        }
    }
    fmt.Println(string(merged))
    return string(merged)
}
