import type { PageData } from "./types";
import axios from "axios";

export interface Project {
    id: number;
    name: string;
    description: string;
    state: number;
    createdBy: number;
    createdAt: string;
    updatedAt: string;
}

export function listProjects(page: number = 1) {
    return axios.get<PageData<Project>>("/api/project");
}

export function createProject() {}
