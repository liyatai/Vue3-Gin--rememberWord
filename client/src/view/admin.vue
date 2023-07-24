<template>
    <div class="base">
      <span class="head">单词表管理</span><br>
      <!-- 添加 -->
           <el-input v-model="input" placeholder="请输入要添加的表名" class="input top" maxlength="20"></el-input>
           <el-button type="primary" size="default" @click="add" class="add top" round>添加</el-button><br><br>
      <!-- 搜索 -->
           <el-input v-model="serach" placeholder="请输入要搜索的表"  class="serach top" ></el-input>
           <el-button type="warning" size="default" @click="search" round class="sinput top">搜索</el-button>
           <div class="btable">
            <el-table :data="detailList"  stripe style="width: 90%" class="table" >
                <el-table-column prop="name" label="单词组名" width="180"  align="center"/>
                <el-table-column prop="init" label="创建/修改时间" width="180" align="center"/>
                <el-table-column prop="len" label="单词数量" width="180" align="center"/>
                <el-table-column prop="" label="操作" align="center" >
                   <template #default="scope">
                     <el-button type="success" size="default"   @click="editWord(scope)" round>修改</el-button>
                     <el-button type="danger" size="default"   @click="deleteTable(scope)" round>删除</el-button>
                   </template>    
                </el-table-column>
            </el-table>
           </div>
    </div>
</template>
<script setup lang='ts'>
import { ref } from 'vue';
import {detail} from '../model'
import _axios from '../utils/myAxios';
import router from '../router';
// 添加
let input = ref('')
// 搜索
let serach = ref('')
// 初始化
let detailList = ref<detail[]>();
const init = async () =>{
  const resp =await _axios.get("/admin/detail")
    detailList.value = resp.data
}
init()

// 添加按钮点击事件
const add =async ()=>{
  if(confirm("确定要执行此操作吗？")) {
  // 用户点击确认按钮后执行的代码
  console.log("用户确认操作");
  const resp = await _axios.get('/admin/addTable',{
       params:{
           name:input.value
       }
  })
  console.log(resp.data);
  location.reload()
} else {
  // 用户点击取消按钮后执行的代码
  console.log("用户取消操作");
}
}
// 搜索事件
const search = async()=>{
    const resp = await _axios.get('/admin/search',{
       params:{
           name:serach.value
       }
    })
    console.log(resp.data);
  detailList.value = resp.data
}
// 删除单词表事件
const deleteTable =async (index:any)=>{
    // 
    if(confirm('确定要执行此操作吗')){
      const params = new URLSearchParams()
      params.append("item",JSON.stringify(detailList.value![index.$index]))
      const resp = await _axios.post('/admin/delTable',params)
      window.alert(resp.data)
      
    }else{
      console.log('用户拒绝了操作');
    }
    location.reload()
}
// 修改单词事件
const editWord = (index:any)=>{
      sessionStorage.setItem("name",detailList.value![index.$index].name)
      router.push('/manage')
}
</script>
<style scoped>
/* 字体 */
@font-face{
    font-family: 'FZSJ-CAOMWDQH';
    src: url('../assets/FZSJ-CAOMWDQH.TTF') format("truetype");
}
.base{
   left: 0px;
   position: relative;
}
.table{
  margin: auto;
  margin-top: 10px;
  border-radius: 20px;
}
.head{
  font-size: 30px;
}
.base{
  font-family: FZSJ-CAOMWDQH;
}
.add{
  position: absolute;
  left: 280px;
}
.input{
   position: absolute;
   width: 200px;
   left: 60px;
}
.serach{
   position: absolute;
   width: 200px;
   right: 140px;
}
.sinput{
   position: absolute;
   right: 60px;
}
.top{
   top: 50px;
}
.btable{
  width: 100%;
  height: 484px;
  overflow: auto;
}
</style>