package sqlc

// CreatePckgToPckgTxParams contains the input parameters for the createPckgToPckgTx function.
//type CreatePckgToPckgTxParams struct {
//	FromPackageID int64           `json:"from_package_id"`
//	ToPackageID   int64           `json:"to_package_id"`
//	Amount        decimal.Decimal `json:"amount"`
//	UomID         int64           `json:"uom_id"`
//	LabTestID     int64           `json:"lab_test_id"`
//}
//
//// CreatePckgToPckgTxResult is the result of the createPckgToPckgTx function.
//type CreatePckgToPckgTxResult struct {
//	PackageAdjustment              PackageAdjustment          `json:"package_adjustment"`
//	FromPackage                    Package                    `json:"from_package"`
//	ToPackage                      Package                    `json:"to_package"`
//	FromPackageAdjEntry            PackageAdjEntry            `json:"from_entry"`
//	ToPackageAdjEntry              PackageAdjEntry            `json:"to_entry"`
//	SourcePackageChildPackageEntry SourcePackagesChildPackage `json:"source_packages_child_package"`
//}
//
//// CreatePckgToPckgTx performs a package adjustment from one package to the other.
//// It creates the package adjustment, add package entries, and update packages' balance within a database transaction
//func (store *SQLStore) CreatePckgToPckgTx(ctx context.Context, arg CreatePckgToPckgTxParams) (CreatePckgToPckgTxResult, error) {
//	var result CreatePckgToPckgTxResult
//
//	println("FromPackageID (p->p): ", arg.FromPackageID)
//	println("ToPackageID (p->p): ", arg.ToPackageID)
//	println("Quantity (p->p): ", arg.Amount.String())
//	println("UomID (p->p): ", arg.UomID)
//	println("LabTestID (p->p): ", arg.LabTestID)
//
//	err := store.execTx(ctx, func(q *Queries) error {
//		var err error
//
//		result.FromPackageAdjEntry, err = q.CreatePackageAdjEntry(ctx, CreatePackageAdjEntryParams{
//			PackageID: arg.FromPackageID,
//			Amount:    decimal.NewFromFloat(-1).Mul(arg.Amount),
//			UomID:     arg.UomID,
//		})
//		if err != nil {
//			return err
//		}
//
//		result.ToPackageAdjEntry, err = q.CreatePackageAdjEntry(ctx, CreatePackageAdjEntryParams{
//			PackageID: arg.ToPackageID,
//			Amount:    arg.Amount,
//			UomID:     arg.UomID,
//		})
//		if err != nil {
//			return err
//		}
//
//		result.SourcePackageChildPackageEntry, err = q.AssignSourcePackageChildPackage(ctx, AssignSourcePackageChildPackageParams{
//			SourcePackageID: arg.FromPackageID,
//			ChildPackageID:  arg.ToPackageID,
//		})
//		if err != nil {
//			return err
//		}
//
//		result.PackageAdjustment, err = q.CreatePackageAdjustment(ctx, CreatePackageAdjustmentParams{
//			FromPackageID: arg.FromPackageID,
//			ToPackageID:   arg.ToPackageID,
//			Amount:        arg.Amount,
//			UomID:         arg.UomID,
//		})
//		if err != nil {
//			return err
//		}
//
//		if arg.LabTestID != 0 {
//			err = q.AssignLabTestToPackage(ctx, AssignLabTestToPackageParams{
//				LabTestID: nulls.NewInt64(arg.LabTestID),
//				PackageID: arg.ToPackageID,
//			})
//			if err != nil {
//				return err
//			}
//		}
//
//		if arg.FromPackageID < arg.ToPackageID {
//			result.FromPackage, result.ToPackage, err = addPckgQty(ctx, q, arg.FromPackageID, decimal.NewFromFloat(-1).Mul(arg.Amount), arg.ToPackageID, arg.Amount)
//		} else {
//			result.ToPackage, result.FromPackage, err = addPckgQty(ctx, q, arg.ToPackageID, arg.Amount, arg.FromPackageID, decimal.NewFromFloat(-1).Mul(arg.Amount))
//		}
//
//		return err
//	})
//
//	return result, err
//}
//
//func addPckgQty(ctx context.Context, q *Queries, fromPckgID int64, fromAmount decimal.Decimal, toPckgID int64, toAmount decimal.Decimal) (fromPckg Package, toPckg Package, err error) {
//	fromPckg, err = q.AddPackageQuantity(ctx, AddPackageQuantityParams{
//		ID:     fromPckgID,
//		Amount: fromAmount,
//	})
//	if err != nil {
//		return fromPckg, toPckg, err
//	}
//
//	toPckg, err = q.AddPackageQuantity(ctx, AddPackageQuantityParams{
//		ID:     toPckgID,
//		Amount: toAmount,
//	})
//	return fromPckg, toPckg, err
//}
//
//func subtractPckgQty(ctx context.Context, q *Queries, fromPckgID int64, fromAmount decimal.Decimal) (fromPckg Package, err error) {
//	fromPckg, err = q.SubtractPackageQuantity(ctx, SubtractPackageQuantityParams{
//		ID:     fromPckgID,
//		Amount: fromAmount,
//	})
//	if err != nil {
//		return fromPckg, err
//	}
//
//	return fromPckg, err
//}
