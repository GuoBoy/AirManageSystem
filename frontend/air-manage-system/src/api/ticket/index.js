import { post, get, postJson } from '../request'

export function getTicketList(callback) {
    get("/ticket/list").then(res => {
        return callback(res)
    }).catch(err => callback(err))
}

export function delFlightInfo(data, callback) {
    post("/flight/del", data).then(res => {
        return callback(res)
    }).catch(err => callback(err))
}

export function updateFlightInfo(data, callback) {
    postJson("/flight/update", data).then(res => {
        return callback(res)
    }).catch(err => callback(err))
}