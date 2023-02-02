/**
* @license
* Copyright 2020 Dynatrace LLC
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*     http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package rules

type Attributes string

var Attributess = struct {
	PgName                Attributes
	PgTag                 Attributes
	ServiceManagementZone Attributes
	ServiceName           Attributes
	ServiceTag            Attributes
	ServiceType           Attributes
}{
	Attributes("PG_NAME"),
	Attributes("PG_TAG"),
	Attributes("SERVICE_MANAGEMENT_ZONE"),
	Attributes("SERVICE_NAME"),
	Attributes("SERVICE_TAG"),
	Attributes("SERVICE_TYPE"),
}

type ServiceTypes string

var ServiceTypess = struct {
	Amp                  ServiceTypes
	Cics                 ServiceTypes
	Cicsinteraction      ServiceTypes
	Customapplication    ServiceTypes
	Database             ServiceTypes
	Enterpriseservicebus ServiceTypes
	External             ServiceTypes
	Ims                  ServiceTypes
	Imsinteraction       ServiceTypes
	Messaging            ServiceTypes
	Method               ServiceTypes
	Mobile               ServiceTypes
	Process              ServiceTypes
	Queueinteraction     ServiceTypes
	Queuelistener        ServiceTypes
	Remotecall           ServiceTypes
	Rmi                  ServiceTypes
	Saasvendor           ServiceTypes
	Webrequest           ServiceTypes
	Webservice           ServiceTypes
	Website              ServiceTypes
	Zosconnect           ServiceTypes
}{
	ServiceTypes("AMP"),
	ServiceTypes("CICS"),
	ServiceTypes("CICSInteraction"),
	ServiceTypes("CustomApplication"),
	ServiceTypes("Database"),
	ServiceTypes("EnterpriseServiceBus"),
	ServiceTypes("External"),
	ServiceTypes("IMS"),
	ServiceTypes("IMSInteraction"),
	ServiceTypes("Messaging"),
	ServiceTypes("Method"),
	ServiceTypes("Mobile"),
	ServiceTypes("Process"),
	ServiceTypes("QueueInteraction"),
	ServiceTypes("QueueListener"),
	ServiceTypes("RemoteCall"),
	ServiceTypes("RMI"),
	ServiceTypes("SaasVendor"),
	ServiceTypes("WebRequest"),
	ServiceTypes("WebService"),
	ServiceTypes("WebSite"),
	ServiceTypes("ZOSConnect"),
}
