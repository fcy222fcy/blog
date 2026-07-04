import request from './request'

export const getLatestQuestion = () => {
  return request.get('/daily-questions/latest')
}

export const getAllPublishedQuestions = () => {
  return request.get('/daily-questions/all')
}

export const getQuestionByDate = (date) => {
  return request.get(`/daily-questions/date/${date}`)
}

export const getPreviousQuestion = (date) => {
  return request.get(`/daily-questions/previous/${date}`)
}

export const getNextQuestion = (date) => {
  return request.get(`/daily-questions/next/${date}`)
}

export const likeQuestion = (id) => {
  return request.post(`/daily-questions/${id}/like`)
}
