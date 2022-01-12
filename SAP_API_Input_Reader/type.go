package sap_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	MaintenanceItem struct {
		MaintenanceItem           string      `json:"document_no"`
		Plant                     string      `json:"deliver_to"`
		OrderQuantity             string      `json:"quantity"`
		PickedQuantity            string      `json:"picked_quantity"`
		NetPriceAmount            string      `json:"price"`
	    Batch                     string      `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string      `json:"document_no"`
		Status               string      `json:"status"`
		DeliverTo            string      `json:"deliver_to"`
		Quantity             string      `json:"quantity"`
		CompletedQuantity    string      `json:"completed_quantity"`
	    PlannedStartDate     string      `json:"planned_start_date"`
	    PlannedValidatedDate string      `json:"planned_validated_date"`
	    ActualStartDate      string      `json:"actual_start_date"`
	    ActualValidatedDate  string      `json:"actual_validated_date"`
	    Batch                string      `json:"batch"`
		Work              struct {
			WorkNo                   string      `json:"work_no"`
			Quantity                 string      `json:"quantity"`
			CompletedQuantity        string      `json:"completed_quantity"`
			ErroredQuantity          string      `json:"errored_quantity"`
			Component                string      `json:"component"`
			PlannedComponentQuantity string      `json:"planned_component_quantity"`
			PlannedStartDate         string      `json:"planned_start_date"`
			PlannedStartTime         string      `json:"planned_start_time"`
			PlannedValidatedDate     string      `json:"planned_validated_date"`
			PlannedValidatedTime     string      `json:"planned_validated_time"`
			ActualStartDate          string      `json:"actual_start_date"`
			ActualStartTime          string      `json:"actual_start_time"`
			ActualValidatedDate      string      `json:"actual_validated_date"`
			ActualValidatedTime      string      `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema                         string `json:"api_schema"`
	Accepter                        []string `json:"accepter"`
	MaterialCode                      string `json:"material_code"`
	Supplier                          string `json:"plant/supplier"`
	Stock                             string `json:"stock"`
	SalesOrderType                    string `json:"document_type"`
	MaintenanceItemNo                 string `json:"document_no"`
	ScheduleLineDeliveryDate          string `json:"planned_date"`
	ValidatedDate                     string `json:"validated_date"`
    Deleted                           bool   `json:"deleted"`
}

type SDC struct {
	ConnectionKey   string `json:"connection_key"`
	Result          bool   `json:"result"`
	RedisKey        string `json:"redis_key"`
	Filepath        string `json:"filepath"`
	MaintenanceItem struct {
		MaintenancePlan              string      `json:"MaintenancePlan"`
		MaintenanceItem              string      `json:"MaintenanceItem"`
		MaintenanceItemDescription   string      `json:"MaintenanceItemDescription"`
		MaintenanceStrategy          string      `json:"MaintenanceStrategy"`
		MaintenancePlanCategory      string      `json:"MaintenancePlanCategory"`
		MaintenancePlanCallObject    string      `json:"MaintenancePlanCallObject"`
		MaintenancePlanItemPosition  string      `json:"MaintenancePlanItemPosition"`
		MaintenanceItemObjectList    string      `json:"MaintenanceItemObjectList"`
		FunctionalLocationLabelName  string      `json:"FunctionalLocationLabelName"`
		Equipment                    string      `json:"Equipment"`
		Assembly                     string      `json:"Assembly"`
		AdditionalDeviceData         string      `json:"AdditionalDeviceData"`
		TaskListType                 string      `json:"TaskListType"`
		TaskListGroup                string      `json:"TaskListGroup"`
		TaskListGroupCounter         string      `json:"TaskListGroupCounter"`
		OperationSystemCondition     string      `json:"OperationSystemCondition"`
		NumberOfTaskListExecutions   string      `json:"NumberOfTaskListExecutions"`
		MaintNotifTskIsAutomlyDtmnd  string      `json:"MaintNotifTskIsAutomlyDtmnd"`
		MaintenancePlanningPlant     string      `json:"MaintenancePlanningPlant"`
		MaintenancePlannerGroup      string      `json:"MaintenancePlannerGroup"`
		MaintenanceOrderType         string      `json:"MaintenanceOrderType"`
		NotificationType             string      `json:"NotificationType"`
		MaintenanceActivityType      string      `json:"MaintenanceActivityType"`
		MainWorkCenter               string      `json:"MainWorkCenter"`
		MainWorkCenterPlant          string      `json:"MainWorkCenterPlant"`
		MaintPriority                string      `json:"MaintPriority"`
		MaintPriorityType            string      `json:"MaintPriorityType"`
		BusinessArea                 string      `json:"BusinessArea"`
		ImmediateReleaseIsBlocked    bool        `json:"ImmediateReleaseIsBlocked"`
		Material                     string      `json:"Material"`
		SerialNumber                 string      `json:"SerialNumber"`
		ServiceDocumentType          string      `json:"ServiceDocumentType"`
		ServiceContract              string      `json:"ServiceContract"`
		ServiceContractItem          string      `json:"ServiceContractItem"`
		ServiceOrderTemplate         string      `json:"ServiceOrderTemplate"`
		ServiceDocumentPriority      string      `json:"ServiceDocumentPriority"`
		Product                      string      `json:"Product"`
		MaintenancePlant             string      `json:"MaintenancePlant"`
		AssetLocation                string      `json:"AssetLocation"`
		AssetRoom                    string      `json:"AssetRoom"`
		PlantSection                 string      `json:"PlantSection"`
		WorkCenter                   string      `json:"WorkCenter"`
		ABCIndicator                 string      `json:"ABCIndicator"`
		MaintObjectFreeDefinedAttrib string      `json:"MaintObjectFreeDefinedAttrib"`
		CompanyCode                  string      `json:"CompanyCode"`
		MasterFixedAsset             string      `json:"MasterFixedAsset"`
		FixedAsset                   string      `json:"FixedAsset"`
		LocAcctAssgmtBusinessArea    string      `json:"LocAcctAssgmtBusinessArea"`
		CostCenter                   string      `json:"CostCenter"`
		ControllingArea              string      `json:"ControllingArea"`
		WBSElement                   string      `json:"WBSElement"`
		SettlementOrder              string      `json:"SettlementOrder"`
		CycleSetSequence             string      `json:"CycleSetSequence"`
		StandingOrderNumber          string      `json:"StandingOrderNumber"`
		CreationDate                 string      `json:"CreationDate"`
		LastChangeDate               string      `json:"LastChangeDate"`
		LastChangeDateTime           string      `json:"LastChangeDateTime"`
		CallObjects                  struct {
			MaintenancePlanCallNumber    int         `json:"MaintenancePlanCallNumber"`
			MaintenanceOrder             string      `json:"MaintenanceOrder"`
			MaintenanceNotification      string      `json:"MaintenanceNotification"`
			ServiceOrder                 string      `json:"ServiceOrder"`
			MaintCallHorizonIsNotReached bool        `json:"MaintCallHorizonIsNotReached"`
			SchedulingStatus             string      `json:"SchedulingStatus"`
			PlannedStartDate             string      `json:"PlannedStartDate"`
		} `json:"CallObjects"`
	} `json:"MaintenanceItem"`
	APISchema         string   `json:"api_schema"`
	Accepter          []string `json:"accepter"`
	MaintenanceItemNo string   `json:"maintenance_item"`
	Deleted           bool     `json:"deleted"`
}
