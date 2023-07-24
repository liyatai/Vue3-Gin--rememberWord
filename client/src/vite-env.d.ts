/// <reference types="vite/client" />
declare module '*.vue' {
    import { ComponentOptions } from 'vue'
    const componentOptions: ComponentOptions
    export default componentOptions
  }
  module.exports = {
    base: '/static/' // 指定基础路径
};