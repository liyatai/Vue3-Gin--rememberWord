import {word} from '../model'
import { reactive } from 'vue'
import _axios from '../utils/myAxios'
// 刷新单词表名信息
function flush(wordList=reactive<word[]>):word{
      const i=Math.floor(Math.random() * (wordList.length))
      const word:word =( wordList as any)[i] 
      return word 
}
// 刷新单独单词表信息
const update= (name:string,list:word[])=>{
      const params = new URLSearchParams()
      params.append("words",JSON.stringify(list))
      params.append("name",name)
      _axios.post("/admin/update",params);
}
export {flush,update}