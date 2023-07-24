<template>
    <div class="base">
        <audio></audio>
        <el-select v-model="value" class="m-2" placeholder="Select" size="large"  @change="listChange">   
            <el-option v-for="item in sel" :key="item" :label="item" :value="item">{{ item }}</el-option>
        </el-select><br>
        <!-- 记单词模块 -->
         <span class="eng">{{ theWord.eng }}</span><br>
         <span class="chi" v-if="visChi">{{ theWord.chi }}</span><br>
         <div class="buttonGroup">
            <el-button class="first button" type="success" round @click="next">切换</el-button>
            <el-button class="button" type="info" round @click="visableChinese">中文</el-button>
            <el-button class="button" type="warning" round @click="play">发音</el-button>
            <el-button class="button" type="danger" round @click="del">移除</el-button>
         </div>
    </div>
</template>
<script setup lang='ts'>
import { reactive, ref} from 'vue';
import _axios from '../utils/myAxios'
import {word} from '../model/'
import {flush} from '../api/func'
// 响应式单词变量
let theWord = ref<word>({
    eng:"英文",
    chi:"中文"
})
// 响应式单词列表变量
let wordList=reactive<word[]>
//获取单词选项
let sel = ref([]); 
let value = ref('请选择');
// 中文 标签条件渲染标志
let visChi = ref(false);
// 播放音频组件的链接
let src = ref('');

   (async function(){
    const list =await _axios.get('/word/list')
    sel.value=list.data 
   })();
   
// 下拉列表监听事件
async function listChange(){
    const resp = await _axios.get("/word/getword",{
        params:{
            item:value.value
        }
    })
   
    wordList = resp.data
    theWord.value=flush(wordList)
    
}

// 切换事件
const next = ()=>{
    theWord.value = flush(wordList)
}
// 显示中文时间
const visableChinese=()=>{
    visChi.value=visChi.value==true?false:true
}
// 播放音频事件
const play = ()=>{
    src.value = "http://dict.youdao.com/dictvoice?audio="+theWord.value.eng+"&type=2"
    let audio=document.querySelector('audio') as HTMLAudioElement;
    audio = new Audio(src.value)
    audio.play();
}
// 删除当前单词
const del = ()=>{
    const del:word = theWord.value
    const list = (wordList as any)
    for(let i=0;i<wordList.length;i++){
         if(list[i].eng==del.eng)
         if(list.length==1){
            theWord.value.eng="英文"
            theWord.value.chi="中文"
            list.splice()
         }else{
            list.splice(i,1)
            theWord.value=flush(wordList)
         }
    }
}
</script>
<style scoped>

.eng,.chi{
   font-size: 30px;
   display: inline-block;
   margin-top: 90px;
   text-align: center;
}
div span:nth-child(1){
    font-size: 33px;
}
.button{
    margin-top: 120px;
    font-size: 30px;
    width: 125px;
    height: 66px;
}

.m-2{
    margin-top: 27px;
    /* margin-left:290px; */
}
.base{
    position: relative;
}
.buttonGroup{
    position: absolute;
    top: 300px;
    left: 120px;
}
</style>