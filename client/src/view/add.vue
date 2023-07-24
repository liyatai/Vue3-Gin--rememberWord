<template>
   <div>
       <el-input v-model="input" placeholder="请输入名称" class="input" maxlength="20"></el-input>
       
   </div>
</template>
<script setup lang='ts'>
import { ref } from 'vue';
let input = ref('')
console.log(input.value);

</script>
<style scoped>
.input{
   margin-top: 30px;
   width: 200px;
}
</style>