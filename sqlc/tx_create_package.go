package sqlc

import (
	"github.com/gobuffalo/nulls"
)

// CreatePackageTxParams contains the input parameters for the createPackageTx function.
type CreatePackageTxParams struct {
	SourcePackageID     nulls.Int64
	CreatePackageParams CreatePackageParams
}

// CreatePackageTxResult contains the output parameters for the createPackageTx function.
type CreatePackageTxResult struct {
	CreatedPackage                 Package                    `json:"created_package"`
	PackageAdjustment              PackageAdjustment          `json:"package_adjustment"`
	FromPackage                    Package                    `json:"from_package"`
	ToPackage                      Package                    `json:"to_package"`
	PackageTag                     PackageTag                 `json:"package_tag"`
	FromPackageAdjEntry            PackageAdjEntry            `json:"from_entry"`
	ToPackageAdjEntry              PackageAdjEntry            `json:"to_entry"`
	SourcePackageChildPackageEntry SourcePackagesChildPackage `json:"source_packages_child_package"`
}

// CreatePackageTx creates a new package, and optionally creates a package adjustment and/or a package to package relationship.
//func (store *SQLStore) CreatePackageTx(ctx context.Context, arg CreatePackageTxParams) (CreatePackageTxResult, error) {
//	var result CreatePackageTxResult
//
//	err := store.execTx(ctx, func(q *Queries) error {
//		// Create the package
//		pkg, err := q.CreatePackage(ctx, &arg.CreatePackageParams)
//		if err != nil {
//			return err
//		}
//		result.CreatedPackage = *pkg
//
//		// Create a package adjustment
//		pkgAdj, err := q.CreatePackageAdjustment(ctx, &CreatePackageAdjustmentParams{
//			FromPackageID: arg.SourcePackageID.Int64,
//			ToPackageID:   pkg.ID,
//			Amount:        arg.CreatePackageParams.Quantity,
//			UomID:         arg.CreatePackageParams.UomID,
//		})
//		if err != nil {
//			return err
//		}
//		result.PackageAdjustment = *pkgAdj
//
//		// Create package adjustment entries
//		fromPckgAdjEntry, err := q.CreatePackageAdjEntry(ctx, &CreatePackageAdjEntryParams{
//			PackageID: arg.SourcePackageID.Int64,
//			Amount:    decimal.NewFromFloat(-1).Mul(arg.CreatePackageParams.Quantity),
//			UomID:     arg.CreatePackageParams.UomID,
//		})
//		if err != nil {
//			return err
//		}
//		result.FromPackageAdjEntry = *fromPckgAdjEntry
//
//		toPckgAdjEntry, err := q.CreatePackageAdjEntry(ctx, &CreatePackageAdjEntryParams{
//			PackageID: pkg.ID,
//			Amount:    arg.CreatePackageParams.Quantity,
//			UomID:     arg.CreatePackageParams.UomID,
//		})
//		if err != nil {
//			return err
//		}
//		result.ToPackageAdjEntry = *toPckgAdjEntry
//
//		sourcePackageChildPackageEntry, err := q.AssignSourcePackageChildPackage(ctx, &AssignSourcePackageChildPackageParams{
//			SourcePackageID: arg.SourcePackageID.Int64,
//			ChildPackageID:  result.CreatedPackage.ID,
//		})
//		if err != nil {
//			return err
//		}
//		result.SourcePackageChildPackageEntry = *sourcePackageChildPackageEntry
//
//		// Make the quantity transfer
//		//if arg.SourcePackageID.Int64 < pkg.ID {
//		//	result.FromPackage, result.ToPackage, err = addPckgQty(ctx, q, arg.SourcePackageID.Int64, decimal.NewFromFloat(-1).Mul(arg.CreatePackageParams.Quantity), pkg.ID, arg.CreatePackageParams.Quantity)
//		//} else {
//		//	result.ToPackage, result.FromPackage, err = addPckgQty(ctx, q, pkg.ID, arg.CreatePackageParams.Quantity, arg.SourcePackageID.Int64, decimal.NewFromFloat(-1).Mul(arg.CreatePackageParams.Quantity))
//		//}
//		//result.FromPackage, result.ToPackage, err = addPckgQty(ctx, q, arg.SourcePackageID.Int64, decimal.NewFromFloat(-1).Mul(arg.CreatePackageParams.Quantity), pkg.ID, arg.CreatePackageParams.Quantity)
//
//		// Subtract the quantity from the source package
//		//result.FromPackage, err = q.SubtractPackageQuantity(ctx, &SubtractPackageQuantityParams{
//		//	ID:     arg.SourcePackageID.Int64,
//		//	Amount: arg.CreatePackageParams.Quantity,
//		//})
//
//		// Get Lab Test connected to the source package
//		labTest, err := q.GetLabTestByPackageID(ctx, arg.SourcePackageID.Int64)
//
//		// Create Lab Test Package Assignment Entry
//		err = q.AssignLabTestToPackage(ctx, &AssignLabTestToPackageParams{
//			LabTestID: labTest.LabTestID,
//			PackageID: pkg.ID,
//		})
//		if err != nil {
//			return err
//		}
//
//		// Update the package_tag row of the id assigned to new package
//		//result.PackageTag, err = q.UpdatePackageTag(ctx, &UpdatePackageTagParams{
//		//	ID:                pkg.TagID.Int64,
//		//	IsAssigned:        nulls.NewBool(true),
//		//	IsActive:          nulls.NewBool(true),
//		//	IsProvisional:     nulls.NewBool(true),
//		//	AssignedPackageID: nulls.NewInt64(pkg.ID),
//		//})
//		//if err != nil {
//		//	return err
//		//}
//
//		return nil
//	})
//	return result, err
//}
