import path from "path";
import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import Icons from "unplugin-icons/vite";
import IconsResolver from "unplugin-icons/resolver";
import Components from "unplugin-vue-components/vite";
import AutoImport from "unplugin-auto-import/vite";
import { ElementPlusResolver } from "unplugin-vue-components/resolvers";

const pathSrc = path.resolve(__dirname, "src");

// https://vitejs.dev/config/
export default defineConfig({
    resolve: {
        alias: {
            "~/": `${pathSrc}/`,
        },
    },
    css: {
        preprocessorOptions: {
            scss: {
                charset: false,
                additionalData: `@use "~/styles/element/index.scss" as *;`,
            },
            css: { charset: false }
        },
    },
    plugins: [
        vue(),
        AutoImport({
            // 自动导入 Vue 相关函数，如：ref, reactive, toRef 等
            imports: ["vue"],
            // 自动导入 Element Plus 相关函数，如：ElMessage, ElMessageBox... (带样式)
            resolvers: [
                ElementPlusResolver(),
                // 自动导入图标组件
                IconsResolver({
                    prefix: "Icon",
                }),
            ],
            dts: "src/auto-imports.d.ts",
        }),
        Components({
            // allow auto load markdown components under `./src/components/`
            extensions: ["vue", "md"],
            // allow auto import and register components used in markdown
            include: [/\.vue$/, /\.vue\?vue/, /\.md$/],
            resolvers: [
                // 自动导入 Element Plus 组件
                ElementPlusResolver({
                    importStyle: "sass",
                }),
                // 自动注册图标组件
                IconsResolver({
                    enabledCollections: ["ep"],
                }),
            ],
            dts: "src/components.d.ts",
        }),
        Icons({
            autoInstall: true,
        }),
    ],
    
});
