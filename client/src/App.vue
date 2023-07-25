<template>
  <div class="main">
       <div class="left">
           <span class="side" @click="side" data-name="word" >记单词</span>
           <span class="side" @click="side" data-name="admin">管理词库</span>
       </div>
       <div class="right">
          <RouterView class="son"></RouterView>
       </div>
  </div>
</template>
<script setup lang="ts">
import router from './router/index'
function side(e:MouseEvent){
    // 样式控制
    const target = e.target as HTMLElement
    const docSide = document.querySelector('.docSide')
    
    docSide?.classList.remove("docSide")
    target.classList.add("docSide")
    // 路由跳转
    const url = target.dataset.name 
    router.push('/'+url)
}

</script>
<style scoped>
.side:nth-child(1){
   border-radius: 50px 0 0 0 ;
}
.son{
  text-align: center;
  
}
.main{
   margin: 0 auto;
   margin-top: 50px;
   width: 1000px;
   height: 600px;
   background-color: pink;
   border-radius: 50px;
}
.left{
  float: left;
  width: 200px;
  height: 600px;
  background-color: yellow;
  border-radius: 50px 0 0 50px;
}
.right{
  float: left;
  width: 800px;
  height: 600px;
  background-color: rgb(207, 238, 207);
  border-radius: 0 50px 50px 0;
}
.side{
  text-align: center;
  font-size: 19px;
  line-height: 90px;
  display: block;
  height: 90px;
  background-color: #1cdad3;
}
.side:hover{
  background-color: #8fe8e5;
}
.docSide{
  background-color: #c6eeed;
}
</style>
