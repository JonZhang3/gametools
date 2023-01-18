import type { HttpResponse, PageData } from "./types";
import request from "./request";

export interface Project {
    id: number;
    name: string;
    description: string;
    state: number;
    createdBy: number;
    createdAt: string;
    updatedAt: string;
}

const ProjectApi = {
    listProjects(page: number = 1, params?: Partial<Project>): Promise<HttpResponse<PageData<Project>>> {
        return request.get("/api/project", {
            page,
            name: params?.name,
            state: params?.state,
        });
    },
    createProject(data: Partial<Project>): Promise<HttpResponse<Project>> {
        return request.post("/api/project", data);
    },
};

export default ProjectApi;
