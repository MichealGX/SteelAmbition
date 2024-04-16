import Vue from 'vue'

export function request(url, options) {
    const baseUrl = Vue.prototype.$baseUrl;
    return new Promise((resolve, reject) => {
        uni.request({
            url: baseUrl + url,
            ...options,
            success: (res) => {
                resolve(res.data);
            },
            fail: (err) => {
                reject(err);
            },
        });
    });
}
// // 设置请求拦截器
// uni.request.interceptors.request.use(config => {
//     // 从本地存储中获取 token
//     const token = uni.getStorageSync('token')
//     if (token) {
//         // 将 token 添加到请求头部
//         config.header.Authorization = token
//     }
//     return config
// }, error => {
//     return Promise.reject(error)
// })
// // 设置响应拦截器
// uni.request.interceptors.response.use(response => {
//     return response
// }, error => {
//     // 判断返回的状态码或错误信息
//     if (error.code === 401) {
//         // token 无效，表示账号被挤下去
//         // 进行相应的处理，如跳转到登录页面
//         uni.clearStorageSync() // 如果返回401说明登录失效则清空本地缓存 进行重新登录
//         uni.showModal({
//             title: '登录失效请重新登录',
//             showCancel: true,
//             success: (confirm) => {
//                 if (confirm){
//                     uni.reLaunch({ url: '/pagesB/login/Login' })
//                 }
//             }
//         })
        
//     }
//     return Promise.reject(error)
// })