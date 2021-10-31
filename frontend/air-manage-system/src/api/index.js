import { post, postJson, get } from './request'
import store from '../store'

const getusers = (callback) => {
  get("/rln/list").then(res => {
    return callback(res)
  }).catch(err => callback(err))
}

const delUser = (data, callback) => {
  postJson("/rln/del", data).then(res => {
    return callback(res)
  }).catch(err => callback(err))
}





const addTicket = (data, callback) => {
  post("/ticket/add", data).then(res => {
    return callback(res)
  }).catch(err => {
    callback(err)
  })
}

const getTicketList = (callback) => {
  get("/ticket/list").then(res => {
    return callback(res)
  }).catch(err => callback(err))
}

const delTicketInfo = (data, callback) => {
  postJson("/ticket/del", data).then(res => {
    return callback(res)
  }).catch(err => callback(err))
}

const updateTicketInfo = (data, callback) => {
  postJson("/ticket/update", data).then(res => {
    return callback(res)
  }).catch(err => callback(err))
}

export {
  getusers,
  delUser,
  addTicket,
  getTicketList,
  delTicketInfo,
  updateTicketInfo,
}
