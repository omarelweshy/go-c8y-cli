/** @type {import('@docusaurus/types').DocusaurusConfig} */

const path = require('path');
const isDev = process.env.NODE_ENV === 'development';

const isDeployPreview =
  process.env.NETLIFY && process.env.CONTEXT === 'deploy-preview';


// Special deployment for staging locales until they get enough translations
// https://app.netlify.com/sites/docusaurus-i18n-staging
// https://docusaurus-i18n-staging.netlify.app/
const isI18nStaging = process.env.I18N_STAGING === 'true';

const allDocHomesPaths = [
  '/docs/',
  '/docs/master/',
];

const baseUrl = `${process.env.BASE_URL || '/'}`;

/** @type {import('@docusaurus/types').DocusaurusConfig} */
(module.exports = {
  title: 'Cumulocity IoT CLI',
  tagline: 'Cumulocity IoT Command Line Interface',
  url: 'https://reubenmiller.github.io',
  baseUrl,
  onBrokenLinks: isDev ? 'warn' : 'throw',
  onBrokenMarkdownLinks: 'warn',
  favicon: 'img/favicon.ico',
  organizationName: 'reubenmiller',
  projectName: 'go-c8y-cli',
  trailingSlash: true,
  markdown: {
    mermaid: true,
  },
  i18n: {
    defaultLocale: 'en',
    locales: isDeployPreview
      ? // Deploy preview: keep it fast!
        ['en']
      : isI18nStaging
      ? // Staging locales: https://docusaurus-i18n-staging.netlify.app/
        ['en']
      : // Production locales
        ['en'],
  },
  themes: [
    '@docusaurus/theme-live-codeblock',
    '@docusaurus/theme-mermaid',
  ],
  plugins: [
    // [
    //   '@docusaurus/plugin-client-redirects',
    //   {
    //     fromExtensions: ['html'],
    //     createRedirects: function (path) {
    //       // redirect to /docs from /docs/introduction,
    //       // as introduction has been made the home doc
    //       if (allDocHomesPaths.includes(path)) {
    //         return [`${path}/introduction`];
    //       }
    //     },
    //     redirects: [
    //       {
    //         from: ['/docs/support', '/docs/next/support'],
    //         to: '/docs/',
    //       },
    //     ],
    //   },
    // ],
    [
      '@docusaurus/plugin-content-blog',
      {
        /**
         * Required for any multi-instance plugin
         */
        id: 'demos-blog',
        /**
         * URL route for the blog section of your site.
         * *DO NOT* include a trailing slash.
         */
        routeBasePath: 'demos',
        /**
         * Path to data on filesystem relative to site dir.
         */
        path: './demos',
      },
    ],
    [
      '@docusaurus/plugin-ideal-image',
      {
        quality: 70,
        max: 1030, // max resized image's size.
        min: 640, // min resized image's size. if original is lower, use that size.
        steps: 2, // the max number of images generated between min and max (inclusive)
      },
    ],
    [
      '@docusaurus/plugin-pwa',
      {
        debug: isDeployPreview,
        offlineModeActivationStrategies: [
          'appInstalled',
          'standalone',
          'queryString',
        ],
        // Warning: Must be set to false otherwise npm run build will fail with 'unknown module false'!
        // swRegister: false,
        swCustom: path.resolve(__dirname, 'src/sw.js'),
        pwaHead: [
          {
            tagName: 'link',
            rel: 'icon',
            href: 'img/logo.png',
          },
          {
            tagName: 'link',
            rel: 'manifest',
            href: `${baseUrl}manifest.json`,
          },
          {
            tagName: 'meta',
            name: 'theme-color',
            content: 'rgb(37, 194, 160)',
          },
          {
            tagName: 'meta',
            name: 'apple-mobile-web-app-capable',
            content: 'yes',
          },
          {
            tagName: 'meta',
            name: 'apple-mobile-web-app-status-bar-style',
            content: '#000',
          },
          {
            tagName: 'link',
            rel: 'apple-touch-icon',
            href: 'img/logo.png',
          },
          {
            tagName: 'link',
            rel: 'mask-icon',
            href: 'img/logo.svg',
            color: 'rgb(62, 204, 94)',
          },
          {
            tagName: 'meta',
            name: 'msapplication-TileImage',
            content: 'img/logo.png',
          },
          {
            tagName: 'meta',
            name: 'msapplication-TileColor',
            content: '#000',
          },
        ],
      },
    ],
  ],
  themeConfig: {
    docs: {
      sidebar: {
        hideable: true
      },
    },
    prism: {
      theme: require('prism-react-renderer/themes/github'),
      darkTheme: require('prism-react-renderer/themes/dracula'),
      additionalLanguages: ['powershell', 'bash', 'markdown'],
    },
    algolia: {
      apiKey: 'f4645553f294c1bacaef04baf61633f0',
      indexName: 'goc8ycli',
      contextualSearch: false,
      searchParameters: {
        facetFilters: [],
        // facetFilters: ["type:lvl1","type:lvl0"],
      },
      appId: 'SKDHZQ8T6H',
    },
    // googleAnalytics: {
    //   trackingID: 'UA-155263011-1',
    //   anonymizeIP: true,
    // },
    announcementBar: {
      id: 'extensions',
      content:
        '📦 go-c8y-cli now supports extensions. Install the latest version (>=2.30.0) and check out the concepts page 🚀',
    },
    navbar: {
      title: 'go-c8y-cli',
      logo: {
        alt: 'go-c8y-cli Logo',
        src: 'img/logo.svg',
      },
      items: [
        {
          type: 'doc',
          docId: 'demo',
          label: 'Docs',
          position: 'left',
        },
        {
          type: 'doc',
          docId: 'cli',
          label: 'API',
          position: 'left',
        },
        {to: 'blog', label: 'Blog', position: 'left'},
        {to: 'demos', label: 'Demos', position: 'left'},
        {
          href: 'https://github.com/reubenmiller/go-c8y-cli',
          className: "header-github-link",
          'aria-label': "GitHub repository",
          position: 'right',
        },
      ],
    },
    footer: {
      style: 'dark',
      links: [
        {
          title: 'Community',
          items: [
            {
              label: 'Blog',
              to: 'blog/',
            },
            {
              label: 'GitHub',
              href: 'https://github.com/reubenmiller/go-c8y-cli',
            },
          ],
        },
      ],
      copyright: `Copyright © ${new Date().getFullYear()} go-c8y-cli. Built with Docusaurus.`,
    },
  },
  scripts: [
    {
      src: 'https://asciinema.org/a/326455.js',
      async: true,
    },
  ],
  presets: [
    [
      '@docusaurus/preset-classic',
      {
        debug: true,
        docs: {
          sidebarPath: require.resolve('./sidebars.js'),
          editUrl: 'https://github.com/reubenmiller/go-c8y-cli/edit/v2/docs/go-c8y-cli/',
        },
        theme: {
          customCss: require.resolve('./src/css/custom.css'),
        },
        googleAnalytics: {
          trackingID: 'UA-155263011-1',
          anonymizeIP: true,
        },
        gtag: {
          trackingID: 'G-64LM52SE7T',
          anonymizeIP: true,
        },
        blog: {
          showReadingTime: true,
          // Please change this to your repo.
          editUrl:
            'https://github.com/reubenmiller/go-c8y-cli/edit/v2/docs/go-c8y-cli/',
        },
      },
    ],
  ],
});
