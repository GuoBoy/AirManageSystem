import { post, postJson, get } from '../request'


export function addTicket(data, callback) {
    postJson("/v1/add/ticket", data).then(res => {
        return callback(res)
    }).catch(err => {
        callback(err)
    })
}

export function getTicketList(data, callback) {
    post("/v1/tickets", data).then(res => {
        return callback(res)
    }).catch(err => callback(err))
}

export function delTicket(data, callback) {
    post("/v1/del/ticket", data).then(res => {
        return callback(res)
    }).catch(err => callback(err))
}
