const axios = require('axios')
const qs = require('qs')

var baseu = ""

if (process.env.NODE_ENV === "development") {
  baseu = 'http://127.0.0.1:23456'
}
const instance = axios.create({
  baseURL: baseu,
  timeout: 1000
})
const get = (url, params) => {
  return new Promise((resolve, reject) => {
    instance.get(url, params).then(({ data }) => {
      resolve(data)
    }).catch((err) => { reject(err) })
  })
}

const post = (url, d) => {
  return new Promise((resolve, reject) => {
    instance.post(url, qs.stringify(d)).then(({ data }) => {
      resolve(data)
    }).catch((err) => { reject(err) })
  })
}

const postJson = (url, d) => {
  return new Promise((resolve, reject) => {
    instance.post(url, JSON.stringify(d)).then(({ data }) => {
      resolve(data)
    }).catch(err => { reject(err) })
  })
}

export {
  get, post, postJson
}
