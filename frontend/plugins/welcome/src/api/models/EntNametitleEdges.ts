/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API Doctor
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntDoctor,
    EntDoctorFromJSON,
    EntDoctorFromJSONTyped,
    EntDoctorToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntNametitleEdges
 */
export interface EntNametitleEdges {
    /**
     * Doctor holds the value of the doctor edge.
     * @type {Array<EntDoctor>}
     * @memberof EntNametitleEdges
     */
    doctor?: Array<EntDoctor>;
}

export function EntNametitleEdgesFromJSON(json: any): EntNametitleEdges {
    return EntNametitleEdgesFromJSONTyped(json, false);
}

export function EntNametitleEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntNametitleEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'doctor': !exists(json, 'doctor') ? undefined : ((json['doctor'] as Array<any>).map(EntDoctorFromJSON)),
    };
}

export function EntNametitleEdgesToJSON(value?: EntNametitleEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'doctor': value.doctor === undefined ? undefined : ((value.doctor as Array<any>).map(EntDoctorToJSON)),
    };
}


