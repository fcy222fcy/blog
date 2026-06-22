<template>
  <div>
    <el-card>
      <template #header>
        <div style="display: flex; justify-content: space-between; align-items: center">
          <span>{{ isEdit ? '编辑文章' : '新建文章' }}</span>
          <el-button @click="$router.push('/articles')">返回列表</el-button>
        </div>
      </template>

      <el-form :model="form" label-width="80px">
        <el-form-item label="标题">
          <el-input v-model="form.title" placeholder="输入文章标题" />
        </el-form-item>
        <el-form-item label="分类">
          <el-select v-model="form.category_id" placeholder="选择分类">
            <el-option v-for="cat in categories" :key="cat.id" :label="cat.name" :value="cat.id" />
          </el-select>
        </el-form-item>
        <el-form-item label="标签">
          <el-select v-model="form.tag_ids" multiple placeholder="选择标签">
            <el-option v-for="tag in tags" :key="tag.id" :label="tag.name" :value="tag.id" />
          </el-select>
        </el-form-item>
        <el-form-item label="摘要">
          <el-input v-model="form.summary" type="textarea" :rows="3" placeholder="文章摘要" />
        </el-form-item>
        <el-form-item label="状态">
          <el-radio-group v-model="form.status">
            <el-radio value="published">发布</el-radio>
            <el-radio value="draft">草稿</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="内容">
          <el-input v-model="form.content" type="textarea" :rows="20" placeholder="Markdown 内容" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSave">保存</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { getArticleDetail, createArticle, updateArticle } from '../../api/article'
import { getCategoryList } from '../../api/category'
import { getTagList } from '../../api/tag'

const route = useRoute()
const router = useRouter()
const isEdit = computed(() => !!route.params.id)

const form = ref({
  title: '',
  content: '',
  summary: '',
  category_id: '',
  tag_ids: [],
  status: 'draft'
})

const categories = ref([])
const tags = ref([])

const loadData = async () => {
  try {
    const [catRes, tagRes] = await Promise.all([getCategoryList(), getTagList()])
    categories.value = catRes.data || []
    tags.value = tagRes.data || []
  } catch (e) { console.error(e) }

  if (isEdit.value) {
    try {
      const res = await getArticleDetail(route.params.id)
      form.value = { ...form.value, ...res.data }
    } catch (e) { console.error(e) }
  }
}

const handleSave = async () => {
  try {
    if (isEdit.value) {
      await updateArticle(route.params.id, form.value)
    } else {
      await createArticle(form.value)
    }
    ElMessage.success('保存成功')
    router.push('/articles')
  } catch (e) { console.error(e) }
}

onMounted(loadData)
</script>
