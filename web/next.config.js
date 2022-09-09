/** @type {import('next').NextConfig} */
const withPWA = require('next-pwa')({
  dest: 'public'
})

const nextConfig = () => {
	return withPWA({
		reactStrictMode: true,
		swcMinify: true,
		env: {
			AWS_CLIENT_ID: process.env.AWS_CLIENT_ID,
			USER_POOL_CLIENT_SECRET: process.env.USER_POOL_CLIENT_SECRET,
			AWS_REGION: process.env.AWS_REGION,
			ENVIRONMENT: process.env.ENVIRONMENT,
		},
		async headers() {
			return [
				{
					// Apply these headers to all routes in your application.
					source: "/:path*",
					headers: [
						{
							key: "X-DNS-Prefetch-Control",
							value: "on",
						},
						{
							key: "Strict-Transport-Security",
							value: "max-age=63072000; includeSubDomains; preload",
						},
						{
							key: "X-XSS-Protection",
							value: "1; mode=block",
						},
						{
							key: "X-Frame-Options",
							value: "SAMEORIGIN",
						},
						{
							key: "X-Content-Type-Options",
							value: "nosniff",
						},
						{
							key: "Referrer-Policy",
							value: "origin-when-cross-origin",
						},
					],
				},
			]
		},
	})
}

module.exports = nextConfig
