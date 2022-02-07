<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
   
      
        <el-form-item label="端口:">
          
            <el-input v-model="formData.port" clearable placeholder="请输入" />
          
        </el-form-item>
      
        <el-form-item label="host:">
          
            <el-input v-model="formData.host" clearable placeholder="请输入" />
          
        </el-form-item>
      
        <el-form-item label="hostinfo:">
          
            <el-input v-model="formData.hostinfo" clearable placeholder="请输入" />
          
        </el-form-item>
      
        <el-form-item label="url:">
          
            <el-input v-model="formData.url" clearable placeholder="请输入" />
          
        </el-form-item>
      
        <el-form-item label="标题:">
          
            <el-input v-model="formData.title" clearable placeholder="请输入" />
          
        </el-form-item>
      
        <el-form-item label="favicons:">
          
            <el-input v-model="formData.favicons" clearable placeholder="请输入" />
          
        </el-form-item>
      
        <el-form-item label="screenshot:">
          
            <el-input v-model="formData.screenshot" clearable placeholder="请输入" />
          
        </el-form-item>
      
        <el-form-item label="products:">
          
            <el-input v-model="formData.products" clearable placeholder="请输入" />
          
        </el-form-item>
      
        <el-form-item label="protocols:">
          
            <el-input v-model="formData.protocols" clearable placeholder="请输入" />
          
        </el-form-item>
      
        <el-form-item label="端口是否开放:">
          
            
              <el-select v-model="formData.alive" placeholder="请选择" style="width:100%" clearable>
                <el-option v-for="(item,key) in intOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            
          
        </el-form-item>
      
        <el-form-item label="banner:">
          
            <el-input v-model="formData.banner" clearable placeholder="请输入" />
          
        </el-form-item>
      
        <el-form-item label="状态:">
          
            <el-input v-model="formData.status" clearable placeholder="请输入" />
          
        </el-form-item>
      
        <el-form-item label="domain_id:">
          
            <el-input v-model="formData.domain_id" clearable placeholder="请输入" />
          
        </el-form-item>
      
        <el-form-item label="path_ids:">
          
            <el-input v-model="formData.path_ids" clearable placeholder="请输入" />
          
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
  name: 'PortInfo'
}
</script>

<script setup>
import {
  createPortInfo,
  updatePortInfo,
  findPortInfo
} from '@/api/port_info'

import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')



const formData = ref({
        
            port: "", 
          
            host: "", 
          
            hostinfo: "", 
          
            url: "", 
          
            title: "", 
          
            favicons: "", 
          
            screenshot: "", 
          
            products: [], 
          
            protocols: [], 
          
            alive:  undefined , 
          
            banner: "", 
          
            status: "", 
          
            domain_id: "", 
          
            path_ids: [], 
          
        })

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findPortInfo({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.report_info
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
          res = await createPortInfo(formData.value)
          break
        case 'update':
          res = await updatePortInfo(formData.value)
          break
        default:
          res = await createPortInfo(formData.value)
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
