import { defineStore } from 'pinia'
import { ref } from 'vue'
import { getArticleList, getArticleDetail, getArchives } from '../api/article'

export const useArticleStore = defineStore('article', () => {
  const articles = ref([])
  const currentArticle = ref(null)
  const archives = ref([])
  const total = ref(0)
  const loading = ref(false)

  const fetchArticles = async (params = {}) => {
    loading.value = true
    try {
      const res = await getArticleList(params)
      articles.value = res.data.list || []
      total.value = res.data.total || 0
    } catch (error) {
      console.error('获取文章列表失败:', error)
    } finally {
      loading.value = false
    }
  }

  const fetchArticleDetail = async (slug) => {
    loading.value = true
    try {
      const res = await getArticleDetail(slug)
      currentArticle.value = res.data
    } catch (error) {
      console.error('获取文章详情失败:', error)
    } finally {
      loading.value = false
    }
  }

  const fetchArchives = async () => {
    loading.value = true
    try {
      const res = await getArchives()
      archives.value = res.data || []
    } catch (error) {
      console.error('获取文章归档失败:', error)
    } finally {
      loading.value = false
    }
  }

  return {
    articles,
    currentArticle,
    archives,
    total,
    loading,
    fetchArticles,
    fetchArticleDetail,
    fetchArchives
  }
})
