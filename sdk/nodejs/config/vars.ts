// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

declare var exports: any;
const __config = new pulumi.Config("boundary");

/**
 * The base url of the Boundary API, e.g. "http://127.0.0.1:9200". If not set, it will be read from the "BOUNDARY_ADDR" env
 * var.
 */
export declare const addr: string | undefined;
Object.defineProperty(exports, "addr", {
    get() {
        return __config.get("addr");
    },
    enumerable: true,
});

/**
 * The auth method ID e.g. ampw_1234567890
 */
export declare const authMethodId: string | undefined;
Object.defineProperty(exports, "authMethodId", {
    get() {
        return __config.get("authMethodId");
    },
    enumerable: true,
});

/**
 * The auth method login name for password-style auth methods
 */
export declare const passwordAuthMethodLoginName: string | undefined;
Object.defineProperty(exports, "passwordAuthMethodLoginName", {
    get() {
        return __config.get("passwordAuthMethodLoginName");
    },
    enumerable: true,
});

/**
 * The auth method password for password-style auth methods
 */
export declare const passwordAuthMethodPassword: string | undefined;
Object.defineProperty(exports, "passwordAuthMethodPassword", {
    get() {
        return __config.get("passwordAuthMethodPassword");
    },
    enumerable: true,
});

/**
 * Can be a heredoc string or a path on disk. If set, the string/file will be parsed as HCL and used with the recovery KMS
 * mechanism. While this is set, it will override any other authentication information; the KMS mechanism will always be
 * used. See Boundary's KMS docs for examples: https://boundaryproject.io/docs/configuration/kms
 */
export declare const recoveryKmsHcl: string | undefined;
Object.defineProperty(exports, "recoveryKmsHcl", {
    get() {
        return __config.get("recoveryKmsHcl");
    },
    enumerable: true,
});

/**
 * When set to true, does not validate the Boundary API endpoint certificate
 */
export declare const tlsInsecure: boolean | undefined;
Object.defineProperty(exports, "tlsInsecure", {
    get() {
        return __config.getObject<boolean>("tlsInsecure");
    },
    enumerable: true,
});

/**
 * The Boundary token to use, as a string or path on disk containing just the string. If set, the token read here will be
 * used in place of authenticating with the auth method specified in "auth_method_id", although the recovery KMS mechanism
 * will still override this. Can also be set with the BOUNDARY_TOKEN environment variable.
 */
export declare const token: string | undefined;
Object.defineProperty(exports, "token", {
    get() {
        return __config.get("token");
    },
    enumerable: true,
});
