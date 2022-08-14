import { sveltekit } from '@sveltejs/kit/vite';
import path from 'path';

/** @type {import('vite').UserConfig} */
const config = {
	plugins: [sveltekit()],
	resolve: {
		alias: {
			'@sdk': path.resolve('./src/basin-sdk'),
			'@lib': path.resolve('./src/lib'),
			'@util': path.resolve('./src/util')
		}
	}
};

export default config;
