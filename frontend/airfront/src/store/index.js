import { createStore } from 'vuex'

const state = {
  userinfo: {
    username: "",
    token: "",
  }
}

const mutations = {
  updateToken(state, t) {
    state.userinfo.token = t
    localStorage.setItem("userinfo", JSON.stringify(state.userinfo))
  },
  updateusername(state, n) {
    state.userinfo.username = n
    localStorage.setItem("userinfo", JSON.stringify(state.userinfo))
  },
  updateUserInfo(state, { u, t }) {
    state.userinfo.username = u,
      state.userinfo.token = t
    localStorage.setItem("userinfo", JSON.stringify(state.userinfo))
  }
}

function loadUserInfo(state) {
  try {
    var data = localStorage.getItem("userinfo")
    if (data) {
      state.userinfo = JSON.parse(data)
    }
  } catch (_) { }
}

const getters = {
  getToken(state) {
    if (state.userinfo.token === "") loadUserInfo(state)
    return state.userinfo.token
  },
  getUserName(state) {
    if (state.userinfo.username === "") loadUserInfo(state)
    return state.userinfo.username
  },
}

export default createStore({
  state,
  mutations,
  getters,
  actions: {
  },
  modules: {
  }
})
