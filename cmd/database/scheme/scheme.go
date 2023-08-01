package scheme

import (
	"github.com/bagusyanuar/go-olin-bags-with-fiber/common"
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type User struct {
	ID       uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	Email    string         `gorm:"index:idx_email,unique;type:varchar(255);" json:"email"`
	Username string         `gorm:"index:idx_username,unique;type:varchar(255);not null" json:"username"`
	Password *string        `gorm:"type:text" json:"password"`
	Roles    datatypes.JSON `gorm:"type:longtext;not null" json:"roles"`
	common.WithTimestampsModel
}

type Province struct {
	ID   uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	Code string    `gorm:"column:code;type:char(4);index:idx_code,unique;" json:"code"`
	Name string    `gorm:"column:name;type:varchar(255);" json:"name"`
	common.WithTimestampsModel
}

type City struct {
	ID         uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	ProvinceID uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_province_id;not null" json:"province_id"`
	Code       string    `gorm:"column:code;type:char(4);index:idx_code,unique;" json:"code"`
	Name       string    `gorm:"column:name;type:varchar(255);" json:"name"`
	Province   Province  `gorm:"foreignKey:ProvinceID" json:"province"`
	common.WithTimestampsModel
}

type ProductionHouse struct {
	ID        uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	CityID    uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_city_id;not null" json:"city_id"`
	Name      string    `gorm:"type:varchar(255);not null;" json:"name"`
	Phone     string    `gorm:"type:varchar(25);not null;" json:"phone"`
	Address   string    `gorm:"type:text" json:"address"`
	Latitude  float64   `gorm:"type:decimal(10,8)" json:"latitude"`
	Longitude float64   `gorm:"type:decimal(11,8)" json:"longitude"`
	City      City      `gorm:"foreignKey:CityID" json:"city"`
	common.WithTimestampsModel
}

type Agent struct {
	ID        uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	UserID    uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_user_id;not null;" json:"user_id"`
	CityID    uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_city_id;not null" json:"city_id"`
	Name      string    `gorm:"type:varchar(255);not null;" json:"name"`
	Phone     string    `gorm:"type:varchar(25);not null;" json:"phone"`
	Address   string    `gorm:"type:text" json:"address"`
	Latitude  float64   `gorm:"type:decimal(10,8)" json:"latitude"`
	Longitude float64   `gorm:"type:decimal(11,8)" json:"longitude"`
	Balance   float64   `gorm:"type:double(20,2);default:0;" json:"balance"`
	User      User      `gorm:"foreignKey:UserID" json:"user"`
	City      City      `gorm:"foreignKey:CityID" json:"city"`
	common.WithTimestampsModel
}

type SewingAgent struct {
	ID                uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	ProductionHouseID uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_production_house_id;not null;" json:"production_house_id"`
	Name              string    `gorm:"type:varchar(255);not null;" json:"name"`
	Phone             string    `gorm:"type:varchar(25);not null;" json:"phone"`
	Address           string    `gorm:"type:text" json:"address"`
	common.WithTimestampsModel
	ProductionHouse ProductionHouse `gorm:"foreignKey:ProductionHouseID" json:"production_house"`
}

type PrintingAgent struct {
	ID                uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	ProductionHouseID uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_production_house_id;not null;" json:"production_house_id"`
	Name              string    `gorm:"type:varchar(255);not null;" json:"name"`
	Phone             string    `gorm:"type:varchar(25);not null;" json:"phone"`
	Address           string    `gorm:"type:text" json:"address"`
	common.WithTimestampsModel
	ProductionHouse ProductionHouse `gorm:"foreignKey:ProductionHouseID" json:"production_house"`
}

type Material struct {
	ID   uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	Name string    `gorm:"type:varchar(255);not null;" json:"name"`
	common.WithTimestampsModel
}

type Item struct {
	ID          uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	MaterialID  uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_material_id;not null;" json:"material_id"`
	Name        string    `gorm:"type:varchar(255);not null;" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	Price       uint64    `gorm:"type:bigint(20);default=0" json:"price"`
	common.WithTimestampsModel
	Material Material `gorm:"foreignKey:MaterialID" json:"material"`
}

type Purchasing struct {
	ID                  uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	PurchaserID         uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_purchasher_id;not null;" json:"purchaser_id"`
	ProductionHouseID   uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_production_house_id;not null;" json:"production_house_id"`
	AccessorID          uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_accessor_id;" json:"accessor_id"`
	PurchaseNumber      string         `gorm:"type:varchar(255);not null;unique" json:"purchase_number"`
	Date                datatypes.Date `gorm:"type:date" json:"date"`
	ShippingDestination string         `gorm:"type:text" json:"shipping_destination"`
	SubTotal            uint64         `gorm:"type:bigint(20);default=0" json:"sub_total"`
	ShippingCost        uint64         `gorm:"type:bigint(20);default=0" json:"shipping_cost"`
	Discount            uint64         `gorm:"type:bigint(20);default=0" json:"discount"`
	Total               uint64         `gorm:"type:bigint(20);default=0" json:"total"`
	Status              uint8          `gorm:"type:smallint(6);default=0" json:"status"`
	common.WithTimestampsModel
	Purchaser       User            `gorm:"foreignKey:PurchaserID" json:"purchaser"`
	ProductionHouse ProductionHouse `gorm:"foreignKey:ProductionHouseID" json:"production_house"`
	Accessor        User            `gorm:"foreignKey:AccessorID" json:"accessor"`
}

type PurchaseItem struct {
	ID           uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	PurchasingID uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_purchasing_id;" json:"purchasing_id"`
	PurchaserID  uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_purchasher_id;not null;" json:"purchaser_id"`
	ItemID       uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_item_id;" json:"item_id"`
	Price        uint64    `gorm:"type:bigint(20);default=0" json:"price"`
	Qty          uint32    `gorm:"type:int(11);default=0" json:"qty"`
	Total        uint64    `gorm:"type:bigint(20);default=0" json:"total"`
	common.WithTimestampsModel
	Purchasing Purchasing `gorm:"foreignKey:PurchasingID" json:"Purchasing"`
	Purchaser  User       `gorm:"foreignKey:PurchaserID" json:"purchaser"`
	Item       Item       `gorm:"foreignKey:ItemID" json:"item"`
}
