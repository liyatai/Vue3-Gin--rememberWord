<template>
   <div class="base">
        <span class="head">{{ name }}</span>
         <!-- 添加 -->
           <el-button type="primary" size="default"  class="add top" @click="add" round>添加</el-button><br><br>
      <!-- 搜索 -->
           <el-input v-model="serach" placeholder="请输入要搜索的单词"   class="serach top" ></el-input>
           <el-button type="warning" size="default"  round class="sinput top" @click="find">搜索</el-button>
           <div class="btable">
            <el-table :data="tableList"  stripe style="width: 90%" class="table" >
                <el-table-column type="index" label="序号" width="110" align="center"/>
                <el-table-column prop="eng" label="英文" width="180"  align="center"/>
                <el-table-column prop="chi" label="中文" width="180" align="center"/>
                <el-table-column prop="" label="操作" align="center" >
                   <template #default="scope">
                     <el-button type="success" size="default" @click="Change(scope)"  round>修改</el-button>
                     <el-button type="danger" size="default" @click="del(scope)"  round>删除</el-button>
                   </template>    
                </el-table-column>
            </el-table>
           </div>
   </div>
</template>
<script setup lang='ts'>
import {ref} from 'vue'
import { word } from '../model';
import _axios from '../utils/myAxios';
import {update} from '../api/func'
const name = ref(sessionStorage.getItem("name"))
// 添加

// 搜索
let serach = ref('')
// 初始化
let tableList = ref<word[]>();
(async()=>{
   const resp = await _axios.get('/word/getword',{
      params:{
         item:name.value
      }
   })
   tableList.value=resp.data
   
})()

const Change =async (index:any)=>{
   // 先获取表格值，防止搜索后改变范围增大
   const resp = await _axios.get('/word/getword',{
      params:{
         item:name.value
      }
   })
   const wordList:word[]=resp.data
   // 修改表格的值
   const obj:word = tableList.value![index.$index]
   for(let i=0;i<wordList.length;i++){
      if(wordList[i].eng==tableList.value![index.$index].eng){
           wordList[i].eng= prompt("输入英文",obj.eng) as string
           wordList[i].chi = prompt("输入英文",obj.chi) as string
      }
   }
   update(name.value!,wordList)
}
// 添加单词事件
const add = ()=>{
   let addWord:word = {
       eng:"",
       chi:""
   }
   addWord.eng = prompt("请输入英文")!
   addWord.chi = prompt("请输入中文")!
   if(tableList.value==null){
      tableList.value=[]
   }
   tableList.value?.push(addWord)
   let params = new URLSearchParams()
   params.append("list",JSON.stringify(tableList.value))
   params.append("name",name.value!)
   _axios.post("/admin/addWord",params)
}
// 删除单词事件
const del = (scope:any)=>{
   const word = tableList.value![scope.$index]
   tableList.value?.splice(scope.$index,1)
   _axios.get("/admin/delWord",{
       params:{
          word:JSON.stringify(word),
          name:name.value
       }
   })
}
// 搜索事件
const find =async()=>{
   const resp = await _axios.get("/admin/findWord",{
      params:{
         item:serach.value,
         table:name.value
      }
    })
    tableList.value=resp.data
    console.log(resp.data);
    
}
</script>
<style scoped>
.base{
   position: relative;
}
.head{
   font-size: 30px;
}
.add{
  position: absolute;
  left: 130px;
}
.input{
   position: absolute;
   width: 200px;
   left: 60px;
}
.serach{
   position: absolute;
   width: 200px;
   right: 220px;
}
.sinput{
   position: absolute;
   right: 130px;
}
.top{
   top: 50px;
}
.btable{
  width: 100%;
  height: 484px;
  overflow: auto;
}
.table{
  margin: auto;
  margin-top: 40px;
  border-radius: 20px;
}
</style>