/** @type {import('next').NextConfig} */
const nextConfig = {
  i18n: {
    defaultLocale: 'zh-CN',
    locales: ['en', 'zh-CN'],
  },
  reactStrictMode: true,
  compiler: {
    styledComponents: true,
  }
};

export default nextConfig;
