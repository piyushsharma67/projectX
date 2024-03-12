import axios, { AxiosInstance, AxiosRequestConfig, CancelToken } from "axios";
import { apiUrl } from "../../apiUrl/appUrl";

export default class ServiceApi {
    private _axios: AxiosInstance;

    private static _instance: ServiceApi;

    private constructor(baseUrl: string) {
        this._axios = axios.create({
            baseURL: baseUrl,
            validateStatus: _status => true,
            responseType: "json",
            withCredentials: false
        });
    }

    // use singleton method to create only one instance of the class.
    public static getInstance = () => {
        if (!ServiceApi._instance)
            ServiceApi._instance = new ServiceApi(apiUrl);
        return ServiceApi._instance;
    }

    /* eslint-disable */
    public getCancelTokenSource = () => axios.CancelToken.source();

    public GET = (url: string, config?: AxiosRequestConfig, cancelToken?: CancelToken) => this._axios.get(url, { ...config, cancelToken });

    public POST = (url: string, data: any, config?: AxiosRequestConfig, cancelToken?: CancelToken) => this._axios.post(url, data, { ...config, cancelToken });

    public PUT = (url: string, data: any, config?: AxiosRequestConfig, cancelToken?: CancelToken) => this._axios.put(url, data, { ...config, cancelToken });

    public PATCH = (url: string, data: any, config?: AxiosRequestConfig, cancelToken?: CancelToken) => this._axios.patch(url, data, { ...config, cancelToken });

    public DELETE = (url: string, config?: AxiosRequestConfig, cancelToken?: CancelToken) => this._axios.delete(url, { ...config, cancelToken });
}