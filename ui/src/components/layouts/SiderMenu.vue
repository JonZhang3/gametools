<script lang="tsx">
import type { RouteRecordNormalized, RouteRecordRaw } from "vue-router";
import { useRouter } from "vue-router";
import { compile, computed, defineComponent, h, ref } from "vue";
import { listenerRouteChange } from "@/utils/route-listener";

export default defineComponent({
    setup() {
        const router = useRouter();
        const appRouters = computed(() => {
            return router
                .getRoutes()
                .find((el) => el.name === "root") as RouteRecordNormalized;
        });
        const selectedKeys = ref<string[]>([]);
        const goto = (item: RouteRecordRaw) => {
            router.push({
                name: item.name
            });
        };

        const menuTree = computed(() => {
            const copyRouter = JSON.parse(JSON.stringify(appRouters.value.children));
            copyRouter.sort(
                (a: RouteRecordNormalized, b: RouteRecordNormalized) => {
                    return (a.meta.order as number || 0) - (b.meta.order as number || 0);
                }
            );

            function travel(_routes: RouteRecordRaw[], layer: number) {
                if (!_routes) return null;
                const collector: any = _routes.map((element) => {
                    // leaf node
                    if (!element.children) {
                        return element;
                    }

                    // route filter hideInMenu true
                    element.children = element.children.filter(
                        (x) => x.meta?.hideInMenu !== true
                    );

                    // Associated child node
                    const subItem = travel(element.children, layer);
                    if (subItem.length) {
                        element.children = subItem;
                        return element;
                    }
                    // the else logic
                    if (layer > 1) {
                        element.children = subItem;
                        return element;
                    }

                    if (element.meta?.hideInMenu !== true) {
                        return element;
                    }

                    return null;
                });
                return collector.filter(Boolean);
            }

            return travel(copyRouter, 0);
        });

        listenerRouteChange((newRoute) => {
            const matched = newRoute.matched;
            for (let i = matched.length - 1; i >= 0; i--) {
                const item = matched[i];
                if (item.meta?.hideInMenu !== true) {
                    selectedKeys.value = [item.name as string];
                    break;
                }
            }
        }, true);

        const renderSubMenu = () => {
            function travel(_route: RouteRecordRaw[], nodes = []) {
                if (_route) {
                    _route.forEach(element => {
                        const icon = element?.meta?.icon
                            ? `<${element?.meta?.icon}/>`
                            : ``;
                        const hasChildren = element.children && element.children.length > 0;
                        const r = hasChildren ? (
                            <a-sub-menu key={element?.name} v-slots={{
                                icon: () => h(compile(icon)),
                                title: () => element?.meta?.title
                            }}>
                                {
                                    element.children?.map((elem) => {
                                        return elem?.meta?.menu ? (
                                            <a-menu-item key={elem?.name} onClick={() => goto(elem)}>
                                                {elem?.meta?.title}
                                                {travel(elem.children ?? [])}
                                            </a-menu-item>
                                        ) : null;
                                    })
                                }
                            </a-sub-menu>
                        ) : (
                            <a-menu-item
                                key={element?.name}
                                v-slots={{ icon: () => h(compile(icon)) }}
                                onClick={() => goto(element)}>
                                {element?.meta?.title}
                            </a-menu-item>
                        );
                        nodes.push(r as never);
                    });
                }
                return nodes;
            }

            return travel(menuTree.value);
        };

        return () => {
            <a-menu auto-open={false} selected-keys={selectedKeys.value} auto-open-selected
                    level-indent={34} style="height: 100%">
                {renderSubMenu()}
            </a-menu>;
        };
    }
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
