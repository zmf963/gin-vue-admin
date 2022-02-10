<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
   
      
        <el-form-item label="漏洞名称:">
          
            <el-input v-model="formData.vul_name" clearable placeholder="请输入" />
          
        </el-form-item>
      
        <el-form-item label="漏洞类型:">
          
            <el-input v-model="formData.vul_type" clearable placeholder="请输入" />
          
        </el-form-item>
      
        <el-form-item label="是否登录:">
          
            <el-switch v-model="formData.is_login" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
          
        </el-form-item>
      
        <el-form-item label="漏洞编号:">
          
            <el-input v-model="formData.vul_id" clearable placeholder="请输入" />
          
        </el-form-item>
      
        <el-form-item label="漏洞厂商:">
          
            <el-input v-model="formData.vul_manufacturer" clearable placeholder="请输入" />
          
        </el-form-item>
      
        <el-form-item label="漏洞系统:">
          
            <el-input v-model="formData.vul_system" clearable placeholder="请输入" />
          
        </el-form-item>
      
        <el-form-item label="系统语言:">
          
            <el-input v-model="formData.language" clearable placeholder="请输入" />
          
        </el-form-item>
      
        <el-form-item label="系统版本:">
          
            <el-input v-model="formData.version" clearable placeholder="请输入" />
          
        </el-form-item>
      
        <el-form-item label="相关文章:">
          
            <el-input v-model="formData.link_url" clearable placeholder="请输入" />
          
        </el-form-item>
      
        <el-form-item label="poc类型:">
          
            <el-input v-model="formData.poc_type" clearable placeholder="请输入" />
          
        </el-form-item>
      
        <el-form-item label="poc内容:">
          
            <el-input v-model="formData.poc_content" clearable placeholder="请输入" />
          
        </el-form-item>
      
        <el-form-item label="poc参数:">
          
            <el-input v-model="formData.poc_args" clearable placeholder="请输入" />
          
        </el-form-item>
      
        <el-form-item label="漏洞指纹id:">
          
            <el-input v-model="formData.vul_finger_id" clearable placeholder="请输入" />
          
        </el-form-item>
      

        <el-form-item>
          <el-button size="mini" type="primary" @click="save">保存</el-button>
          <el-button size="mini" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'PocInfo'
}
</script>

<script setup>
import {
  createPocInfo,
  updatePocInfo,
  findPocInfo
} from '@/api/poc_info'

import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')



const formData = ref({
        
            vul_name: "", 
          
            vul_type: "", 
          
            is_login: false, 
          
            vul_id: "", 
          
            vul_manufacturer: "", 
          
            vul_system: "", 
          
            language: "", 
          
            version: "", 
          
            link_url: [], 
          
            poc_type: "", 
          
            poc_content: "", 
          
            poc_args: [], 
          
            vul_finger_id: "", 
          
        })

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findPocInfo({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.repoc_info
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }

}

init()
// 保存按钮
const save = async() => {
      let res
      switch (type.value) {
        case 'create':
          res = await createPocInfo(formData.value)
          break
        case 'update':
          res = await updatePocInfo(formData.value)
          break
        default:
          res = await createPocInfo(formData.value)
          break
      }
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '创建/更改成功'
        })
      }
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
