// svelte.config.js
import adapter from '@sveltejs/adapter-static';

export default {
	kit: {
		adapter: adapter({
			// default options are shown
			pages: 'build',
			assets: 'build',
			fallback: null
		})
	},
	onwarn: (warning, handler) => {
        const { code, frame } = warning;
        if (code === "css-unused-selector")
            return;

        handler(warning);
    },
};