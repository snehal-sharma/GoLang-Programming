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
    return string(merged)
}

func mergeAlternately2(word1 string, word2 string) string {
    merged := ""
    for i:=0;i<len(word1) || i<len(word2);i++{
        str1,str2:="",""
        if i<len(word1){
            str1=string(word1[i])
        }
        if i<len(word2){
            str2=string(word2[i])
        }
        merged+=str1+str2
    }
    fmt.Println(string(merged))
    return string(merged)
}
