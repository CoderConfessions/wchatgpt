import type { Message } from "./conversation";

export function useGPT() {
  const send = async (message: string) => {
		
		try {
			const resp = await wx.cloud.callContainer({
				"config": {
					"env": "prod-0g1idfbhf182d0c2"
				},
				"path": "/chat/text_completion",
				"header": {
					"X-WX-SERVICE": "golang-1jmj",
					"Content-Type": "application/json"
				},
				"method": "POST",
				"data": {
					"action": "inc",
					"prompt": message
				}
			});
			console.log(resp.data);
		}catch (error) {
			console.error(error);
		}
		return message;
	}
  return { send };
}
