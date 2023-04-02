import type { Message } from "./conversation";

export function useGPT() {	
	const send = async (message: string) => {
		const resp = await new Promise((resolve, reject) => {
			wx.request({
				url: 'http://localhost:8080/chat/text_completion',
				data: {
					prompt: message,
					user_uid: '0000'
				},
				method: 'POST',
				header: {
					'Content-Type': 'application/json'
				},
				success(res) {
					// 成功处理返回数据
					console.log(res.data);
					resolve(res.data);
				},
				fail(error) {
					// 处理请求失败
					console.error(error);
					reject(error);
				},
				complete() {
					// 请求完成时回调，成功或失败都会执行
					console.log("complete");
				}
			});
		});
		return resp;
	}
	return send ;
}
