<script lang="tsx">
import type { RouteRecordNormalized, RouteRecordRaw } from "vue-router";
import { useRouter } from "vue-router";
import { compile, computed, defineComponent, h, ref } from "vue";
import { listenerRouteChange } from "@/utils/route-listener";

export default defineComponent({
    props: {
        name: {
            type: String,
            default: "root",
        },
    },
    setup(props) {
        const name = ref(props.name);
        const router = useRouter();
        const appRouters = computed(() => {
            return router.getRoutes().filter((el) => el.name === name.value);
        });
        const selectedKeys = ref<string[]>([]);
        const goto = (item: RouteRecordRaw) => {
            router.push({
                name: item.name,
            });
        };

        listenerRouteChange((newRoute) => {
            const matched = newRoute.matched;
            for (let i = matched.length - 1; i >= 0; i--) {
                const item = matched[i];
                if (item.meta?.title) {
                    selectedKeys.value = [item.name as string];
                    break;
                }
            }
        }, true);

        const renderSubMenu = () => {
            function travel(_route: RouteRecordNormalized[], nodes = []) {
                if (_route) {
                    _route.forEach((el) => {
                        const isLayout = el.meta?.layout;
                        const hasChildren = el.children && el.children.length > 0;
                        if (!isLayout && !el.meta?.title) {
                            return;
                        }
                        const icon = el.meta?.icon ? `<${el.meta.icon}/>` : "";
                        const r = hasChildren ? (
                            <a-sub-menu
                                key={el.name}
                                v-slots={{
                                    icon: () => h(compile(icon)),
                                    title: () => el.meta?.title,
                                }}
                            >
                                {el.children.map((ch) => {
                                    return ch.meta?.title ? (
                                        <a-menu-item key={ch.name} onClick={() => goto(ch)}>
                                            {ch.meta.title}
                                        </a-menu-item>
                                    ) : null;
                                })}
                            </a-sub-menu>
                        ) : (
                            <a-menu-item
                                key={el.name}
                                v-slots={{ icon: () => h(compile(icon)) }}
                                onClick={() => goto(el)}
                            >
                                {el.meta?.title}
                            </a-menu-item>
                        );
                        nodes.push(r as never);
                    });
                }
                return nodes;
            }
            return travel(appRouters.value);
        };

        return () => (
            <a-menu
                auto-open={false}
                selected-keys={selectedKeys.value}
                auto-open-selected
                level-indent={34}
                style="height: 100%;"
            >
                {renderSubMenu()}
            </a-menu>
        );
    },
});
</script>

<style lang="less" scoped>
:deep(.arco-menu-inner) {
    .arco-menu-inline-header {
        display: flex;
        align-items: center;
    }

    .arco-icon {
        &:not(.arco-icon-down) {
            font-size: 18px;
        }
    }
}
</style>
