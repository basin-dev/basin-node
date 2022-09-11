import { useEffect, useState } from 'react';
import type { BasinSDK } from '.';

export const basinReadHook = (basin: BasinSDK, url: string) => {
	const [resp, setResp] = useState<any>('');
	const [err, _] = useState<any>('');

	useEffect(() => {
		(async () => {
			let resp = await basin.read(url);
			setResp(resp);
		})();
	}, []);

	return { resp, err };
};

export const basinReadMultiHook = (basin: BasinSDK, urls: string[]) => {
	const [resp, setResp] = useState<any>('');
	const [err, _] = useState<any>('');

	useEffect(() => {
		async () => {
			let responses = await Promise.all(
				urls.map(async (url: string) => {
					return await basin.read(url);
				})
			);
			let respObject: any = {};
			for (let i = 0; i < urls.length; i++) {
				respObject[urls[i]] = responses[i];
			}
			setResp(respObject);
		};
	}, []);

	return { resp, err };
};
